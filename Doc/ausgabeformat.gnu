reset
set xrange [<xmin>:<xmax>]
set yrange [<ymin>:<ymax>]
set size ratio 1.0
set title "<titel>, Iteration: <nr>"
unset xtics
unset ytics
$data << EOD
<Liste aus <xpos> <ypos> <radius> <ID> <numericID>>
EOD
plot \
'$data' using 1:2:3:5 with circles lc var notitle, \
'$data' using 1:2:4:5 with labels font "arial,9" tc var notitle
