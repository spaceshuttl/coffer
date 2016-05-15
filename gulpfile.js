'use strict'
const gulp        = require('gulp');
const less        = require('gulp-less');
const react       = require('gulp-react');
const cleanCSS    = require('gulp-clean-css');
const uglify      = require('gulp-uglify')
const path        = require('path');
const browserify  = require('browserify');
const babelify    = require('babelify');
const source      = require('vinyl-source-stream');

gulp.task('less', function () {
  return gulp.src(__dirname + '/src/app/src/less/kube.less')
    .pipe(less())
    .pipe(cleanCSS())
    .pipe(gulp.dest(__dirname + '/dist/coffer/css'))
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
  .pipe(gulp.dest(__dirname + '/dist/coffer/js'));
});


gulp.task('default', ['less', 'react'], function () {
  gulp.watch("./src/app/src/less/**/*.less", ['less'])
  gulp.watch("./src/app/src/**/*.jsx", ['react'])
})

gulp.task('build', ['less', 'react'], function() {
  gulp.src(__dirname + '/src/app/main.js')
  .pipe(gulp.dest('dist/coffer'))

  gulp.src(__dirname + '/src/app/index.html')
  .pipe(gulp.dest('dist/coffer'))
  return
})
