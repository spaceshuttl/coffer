'use strict'

const electron = require('electron')
const request = require('request');
const child_process = require('child_process')
const app = electron.app  // Module to control application life.
const BrowserWindow = electron.BrowserWindow  // Module to create native browser window.

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
var mainWindow = null

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  app.quit()
})

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
app.on('ready', function() {

  // start the backend
  // var service = child_process.execFile("./bin/service", {
  //   env: {
  //     "LEVEL": "info"
  //   }
  // }, (error, stdout, stderr) => {
  //   if(error) {
  //     console.log(error)
  //     app.quit
  //   }
  // })


  // Create the browser window.
  mainWindow = new BrowserWindow({
    title: "Coffer",
    width: 800,
    height: 600,
    frame: false,
    center: true,
  })

  // and load the index.html of the app.key
  mainWindow.loadURL('file://' + __dirname + '/dist/index.html')

  // Open the DevTools.

  if (process.env.DEBUG === "true") {
    mainWindow.webContents.openDevTools()
  }

  // Emitted when the window is closed.
  mainWindow.on('closed', function() {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null
  })
})
