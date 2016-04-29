const gulp = require('gulp');
const less = require('gulp-less');
const path = require('path');
const browserSync = require('browser-sync').create();
const react = require('gulp-react');
const browserify = require('browserify');
const babelify = require('babelify');
const source = require('vinyl-source-stream');

gulp.task('less', function () {
  return gulp.src(__dirname + '/src/app/src/less/kube.less')
    .pipe(less())
    .pipe(gulp.dest(__dirname + '/src/app/dist/css'))
});

gulp.task('react', function () {

  browserify({
    entries: __dirname + '/src/app/src/index.jsx',
    extensions: ['.jsx'],
    debug: true
  })
  .transform(babelify, {presets: ["es2015", "react"]})
  .bundle()
  .pipe(source('bundle.js'))
  .pipe(gulp.dest(__dirname + '/src/app/dist'));
});


gulp.task('default', ['less', 'react'], function () {
  gulp.watch("./src/app/src/less/**/*.less", ['less'])
  gulp.watch("./src/app/src/**/*.jsx", ['react']).on('change', browserSync.reload)
})
