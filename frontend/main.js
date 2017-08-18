'use strict'
const electron      = require('electron')
const {BrowserWindow, app} = require('electron')
const {execFile} = require('child_process')
const os            = require('os')
const fs            = require('fs')
const path          = require('path')
const cwd  = path.join(__dirname, '..')

var win = null

// Quit when all windows are closed.
app.on('window-all-closed', function() {
  app.quit()
})

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
app.on('ready', function() {
  let binName = ""

  // get the operating system
  if (os.type() == 'Linux') {
    binName = "service"
  } else if (os.type() == 'Darwin') {
    binName = "service"
  } else if (os.type() == 'Windows_NT') {
    binName = "service.exe"
  }

  // start the backend
  var service = execFile(`${__dirname}/${binName}`, {
    env: {
      "LEVEL": "debug"
    }
  }, (error, stdout, stderr) => {
    if(error) {
      console.error(error)
      app.quit
    }
  })


  // Create the browser window.
  win = new BrowserWindow({
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

  // Remove the frontend menu
  win.setMenuBarVisibility(false)

  // load the application page
  win.loadURL(`file://${__dirname}/index.html`);

  // Open the DevTools.

  if (process.env.DEBUG === "true") {
    win.webContents.openDevTools()
  }

})
