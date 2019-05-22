reset
set xrange [-33.86357381155014:85.0741647360979]
set yrange [-2.1735154975148223:116.76422305013321]
set size ratio 1.0
set title "Bierkonsum, Iteration: 1000"
unset xtics
unset ytics
$data << EOD
16.751229415522495 63.89562847551619 29.676315947556564 D 0
-23.041037681301255 71.42772638479775 10.822536130248883 NL 1
-20.220018532838186 51.91242303158999 8.895608153727649 B 2
-10.217085354402021 46.65785245565221 2.8470501736687086 L 3
-11.701867227882616 29.77460395147963 14.506691505134297 F 4
9.429475402706583 28.287342791648207 6.676924501791871 CH 5
25.645069779242753 25.449506312958032 9.78511719613052 A 6
46.450989068634634 33.77447230655958 12.624508777350401 CZ 7
65.64943596362681 59.437154636063894 19.42472877247108 PL 8
-3.546190833310066 93.28328965373461 6.039505452538444 DK 9
EOD
plot \
'$data' using 1:2:3:5 with circles lc var notitle, \
'$data' using 1:2:4:5 with labels font "arial,9" tc var notitle
pause -1
