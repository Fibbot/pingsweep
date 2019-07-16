been finding that i've needed to discover hosts on a subnet with machines that may or may not have typical enumeration tools already installed. 
especially cloud hosted networks with a large subnet full of hosts that change ip's daily.

installation:
- clone this into your go src folder
- run compile.sh to compile for windows 32/64, linux 32/64, mac 64
- can edit compile.sh to only compile for what you want

running:
- ./binary 10.10.10.1 24
- runs a scan on 10.10.10.1/24

note: due to using raw sockets, after compiling you will need to run binaries w/ sudo/root on mac/nix... not necessary with Windows


to do:
- fastping has the ability to use UDP as well, haven't added that in yet