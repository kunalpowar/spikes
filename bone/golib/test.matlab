%% green: function description
function [bin_green, num] = green(im)
	[m,n,t] = size(im);
	bin_green = zeroes(m,n);
	num = 0;
	for i=1:m
		for j=1:n
			if(im(i,j,1)==0 && im(i,j,2)==128 && im(i,j,3)==1)
				bin_green(i,j) = 1;
				num = num + 1;
	outputs = ;bin_green, num

