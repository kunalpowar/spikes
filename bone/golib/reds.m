function red_im = red(im)
	[m,n,t] = size(im);
	red_im = zeros(m,n);
	for i=1:m
		for j=1:n
			if(im(i,j,1)>250 && im(i,j,2) <10, im(i,j,3) <10)
				bw(i,j) = [254,0,0]
			end
		end
	end
end