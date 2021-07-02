// +build linux freebsd

package mount

import (
	"bazil.org/fuse"
	fusefs "bazil.org/fuse/fs"
	"context"
	"github.com/ThePolarNight/rclone/fs/log"
	"github.com/ThePolarNight/rclone/vfs"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
)

// File represents a file
type File struct {
	*vfs.File
	fsys *FS
}

// Check interface satisfied
var _ fusefs.Node = (*File)(nil)

// Attr fills out the attributes for the file
func (f *File) Attr(ctx context.Context, a *fuse.Attr) (err error) {
	defer log.Trace(f, "")("a=%+v, err=%v", a, &err)
	a.Valid = f.fsys.opt.AttrTimeout
	modTime := f.File.ModTime()
	Size := uint64(f.File.Size())
	Blocks := (Size + 511) / 512
	a.Gid = f.VFS().Opt.GID
	a.Uid = f.VFS().Opt.UID
	a.Mode = f.VFS().Opt.FilePerms
	if strings.HasSuffix(f.File.Name(), ".rclonelink") {
		defer log.Trace(f, f.File.Name()+" is a symlink")("a=%+v, err=%v", a, &err)
		a.Mode = 0777 | os.ModeSymlink
	}
	a.Size = Size
	a.Atime = modTime
	a.Mtime = modTime
	a.Ctime = modTime
	a.Crtime = modTime
	a.Blocks = Blocks
	return nil
}

// Check interface satisfied
var _ fusefs.NodeSetattrer = (*File)(nil)

// Setattr handles attribute changes from FUSE. Currently supports ModTime and Size only
func (f *File) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) (err error) {
	defer log.Trace(f, "a=%+v", req)("err=%v", &err)
	if !f.VFS().Opt.NoModTime {
		if req.Valid.Mtime() {
			err = f.File.SetModTime(req.Mtime)
		} else if req.Valid.MtimeNow() {
			err = f.File.SetModTime(time.Now())
		}
	}
	if req.Valid.Size() {
		err = f.File.Truncate(int64(req.Size))
	}
	return translateError(err)
}

// Check interface satisfied
var _ fusefs.NodeOpener = (*File)(nil)

// Open the file for read or write
func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fh fusefs.Handle, err error) {
	defer log.Trace(f, "flags=%v", req.Flags)("fh=%v, err=%v", &fh, &err)

	// fuse flags are based off syscall flags as are os flags, so
	// should be compatible
	handle, err := f.File.Open(int(req.Flags))
	if err != nil {
		return nil, translateError(err)
	}

	// If size unknown then use direct io to read
	if entry := handle.Node().DirEntry(); entry != nil && entry.Size() < 0 {
		resp.Flags |= fuse.OpenDirectIO
	}

	return &FileHandle{handle}, nil
}

// Check interface satisfied
var _ fusefs.NodeFsyncer = (*File)(nil)

// Fsync the file
//
// Note that we don't do anything except return OK
func (f *File) Fsync(ctx context.Context, req *fuse.FsyncRequest) (err error) {
	defer log.Trace(f, "")("err=%v", &err)
	return nil
}

// Getxattr gets an extended attribute by the given name from the
// node.
//
// If there is no xattr by that name, returns fuse.ErrNoXattr.
func (f *File) Getxattr(ctx context.Context, req *fuse.GetxattrRequest, resp *fuse.GetxattrResponse) error {
	return fuse.ENOSYS // we never implement this
}

var _ fusefs.NodeGetxattrer = (*File)(nil)

// Listxattr lists the extended attributes recorded for the node.
func (f *File) Listxattr(ctx context.Context, req *fuse.ListxattrRequest, resp *fuse.ListxattrResponse) error {
	return fuse.ENOSYS // we never implement this
}

var _ fusefs.NodeListxattrer = (*File)(nil)

// Setxattr sets an extended attribute with the given name and
// value for the node.
func (f *File) Setxattr(ctx context.Context, req *fuse.SetxattrRequest) error {
	return fuse.ENOSYS // we never implement this
}

var _ fusefs.NodeSetxattrer = (*File)(nil)

// Removexattr removes an extended attribute for the name.
//
// If there is no xattr by that name, returns fuse.ErrNoXattr.
func (f *File) Removexattr(ctx context.Context, req *fuse.RemovexattrRequest) error {
	return fuse.ENOSYS // we never implement this
}

var _ fusefs.NodeRemovexattrer = (*File)(nil)
var _ fusefs.NodeReadlinker = (*File)(nil)

// read symbolic link target
func (f *File) Readlink(ctx context.Context, req *fuse.ReadlinkRequest) (ret string, err error) {
	defer log.Trace(f, "Requested to read link")("ret=%v, err=%v", &ret, &err)

	handle, err := f.File.Open(syscall.O_RDONLY)
	if err != nil {
		return "", translateError(err)
	}
	data := make([]byte, f.File.Size())
	_, err = handle.Read(data)
	if err == io.EOF {
		err = nil
	} else if err != nil {
		return "", translateError(err)
	}

	return string(data), nil
}

// create symbolic link
func (f *File) Symlink(ctx context.Context, req *fuse.SymlinkRequest) (node fusefs.Node, err error) {
	defer log.Trace(f, "target=%q; new_name=%q", req.Target, req.NewName)("node=%v, err=%v", &node, &err)

	symlinkName := req.NewName + ".rclonelink"
	d := f.Dir()
	file, err := d.File.Create(symlinkName, os.O_RDWR)
	if err != nil {
		defer log.Trace(f, "failed to create symlink file "+symlinkName)
		return nil, translateError(err)
	}
	fh, err := file.Open(os.O_RDWR | os.O_CREATE)
	if err != nil {
		defer log.Trace(f, "failed to open symlink file "+symlinkName)
		return nil, translateError(err)
	}
	node = &File{file, f.fsys}

	_, err2 := fh.WriteAt([]byte(req.Target), 0)
	if err2 != nil {
		defer log.Trace(f, "failed to write to symlink file "+symlinkName)
		return nil, translateError(err2)
	}

	err = fh.Release()
	if err != nil {
		defer log.Trace(f, "failed to close symlink file "+symlinkName)
		return nil, translateError(err)
	}

	return node, err
}
