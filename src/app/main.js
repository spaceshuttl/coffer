'use strict'
const electron      = require('electron')
const child_process = require('child_process')
const os            = require('os')
const {BrowserWindow, app} = require('electron')

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
var win = null

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  app.quit()
})

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
app.on('ready', function() {
  // var platform = ""
  // // get our operating system
  // if (os.type() == 'Linux') {
  //   platform = "/service"
  // } else if (os.type() == 'Darwin') {
  //   platform = "/service"
  // } else if (os.type() == 'Windows_NT') {
  //   platform = "/service.exe"
  // }
  //
  // // start the backend
  // var service = child_process.execFile(__dirname + platform, {
  //   env: {
  //     "LEVEL": "debug"
  //   }
  // }, (error, stdout, stderr) => {
  //   if(error) {
  //     console.log(error)
  //     app.quit
  //   }
  // })


  // Create the browser window.
  let win = new BrowserWindow({
    title: "Coffer",
    width: 800,
    height: 600,
    frame: true,
    center: true,
  })

  win.on('closed', function() {
    win = null
  })

  win.once('ready-to-show', () => {
    win.show()
  })

  // Remove the app menu
  win.setMenu(null)

  // load the application page
  win.loadURL(`file://${__dirname}/index.html`);

  // Open the DevTools.

  if (process.env.DEBUG === "true") {
    win.webContents.openDevTools()
  }

})
