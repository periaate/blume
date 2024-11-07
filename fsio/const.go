package fsio

import "os"

// Constants for common file permissions
const (
	// User permissions
	UserRead      = 0o400 // Owner read permission
	UserWrite     = 0o200 // Owner write permission
	UserExecute   = 0o100 // Owner execute permission
	UserReadWrite = UserRead | UserWrite

	// Group permissions
	GroupRead      = 0o040 // Group read permission
	GroupWrite     = 0o020 // Group write permission
	GroupExecute   = 0o010 // Group execute permission
	GroupReadWrite = GroupRead | GroupWrite

	// Other permissions
	OtherRead      = 0o004 // Others read permission
	OtherWrite     = 0o002 // Others write permission
	OtherExecute   = 0o001 // Others execute permission
	OtherReadWrite = OtherRead | OtherWrite

	// Combined permissions
	AllRead      = UserRead | GroupRead | OtherRead          // Read permission for all
	AllWrite     = UserWrite | GroupWrite | OtherWrite       // Write permission for all
	AllExecute   = UserExecute | GroupExecute | OtherExecute // Execute permission for all
	AllReadWrite = AllRead | AllWrite

	// Common file modes
	ReadOnly      = UserRead | GroupRead | OtherRead
	ReadWrite     = UserReadWrite | GroupRead | OtherRead
	ReadWriteExec = ReadWrite | UserExecute | GroupExecute | OtherExecute

	// Directory modes
	DirReadOnly      = ReadOnly | os.ModeDir
	DirReadWrite     = ReadWrite | os.ModeDir
	DirReadWriteExec = ReadWriteExec | os.ModeDir
)
