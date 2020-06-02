set terminal qt
set title 'Dummy 2H 2V'
set xlabel '#lines'
set ylabel 'FPS'
set style fill solid 0.3
set style data lines
plot './results/parsed/bmk-2-2-prof.txt'
