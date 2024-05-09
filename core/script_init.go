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
var copyFile = Y.CopyFile;
var copyDir = Y.CopyDir;
var exportRegistry = Y.ExportRegistry;
var exportReg = Y.ExportRegistry;

// System
var addHostsFile = Y.AddHostsFile;
var exportRegSystemEnv = Y.ExportRegSystemEnv;
var exportRegUserEnv = Y.ExportRegUserEnv;
`
