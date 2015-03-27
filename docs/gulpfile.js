var gulp = require('gulp'),
    shell = require('gulp-shell'),
    watch = require('gulp-watch'),
    plumber = require('gulp-plumber'),


files = ["tex/**/*.tex"] // ,"pdf/**/*.aux","pdf/**/*.log"];

gulp.task('texfiles', function () {
	return gulp.src(files, {read: false})
		.pipe(watch(files))
		.pipe(plumber())
    		.pipe(shell(['pdflatex -halt-on-error -file-line-error --output-directory=pdf  <%= file.path %> ']) )
					
});



gulp.task("default",['texfiles']);
