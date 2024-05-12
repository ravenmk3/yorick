package core

var InitScript = `
// Common
var format = Y.Format;
var logInfo = Y.LogInfo;
var logWarn = Y.LogWarn;
var logError = Y.LogError;
var getEnv = Y.GetEnv;
var taskBegin = Y.TaskBegin;
var taskEnd = Y.TaskEnd;
var destDir = Y.DestDir;
var isDir = Y.IsDir;
var isFile = Y.IsFile;
var fileExt = Y.FileExt;
var listDirs = Y.ListDirs;
var listFiles = Y.ListFiles;
var findLatestFile = Y.FindLatestFile;

var copyFile = Y.CopyFile;
var copyDir = Y.CopyDir;
var copyDirEx = Y.CopyDirEx;
var exportRegistry = Y.ExportRegistry;
var exportReg = Y.ExportRegistry;

var putFile = Y.CopyFile;
var putDir = Y.CopyDir;
var putRegistry = Y.ExportRegistry;
var putReg = Y.ExportRegistry;

// System
var putHostsFile = Y.PutHostsFile;
var exportRegSystemEnv = Y.ExportRegSystemEnv;
var exportRegUserEnv = Y.ExportRegUserEnv;
var putRegSystemEnv = Y.ExportRegSystemEnv;
var putRegUserEnv = Y.ExportRegUserEnv;
`
