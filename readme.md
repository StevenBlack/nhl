# NHL wins minus losses plus/minus standings

This is a personal side-project, written in [Go](https://golang.org/).

I'm unmoved by traditional National Hockey League (NHL) reporting
which [tallies standings by points](https://www.nhl.com/standings/).

Given the disparity, league wide, in games played throughout the season,
total points ranking is always misleading.

The only metric that matters is, wins minus losses.

Generally a team needs 96 points to make the playoffs.
Over an 82-game regular season, 96 points is +14 — a differential of +14,
wins minus losses.

This is what this repo gives you: plaintext plus/minus NHL standings
in the command window.

## Sample output

Show NHL plus/minus team standings in various ways.

* `GP`: games played.
* `+/-`: plus/minus, season so far.
* `+/10`: plus/minus, last 10 games played.
* `GD`: team goal differential, season so far.

``` bash
$ ./nhl

==================================================
NHL Division Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins               40   17    1    34
 2 Tampa Bay Lightning         37    7    3    15
 3 Florida Panthers            38    7    2     8
 4 Toronto Maple Leafs         40    7    7    11
 5 Montréal Canadiens          39    3    2     2
 6 Buffalo Sabres              40    1   -1    -8
 7 Ottawa Senators             40   -3    2   -21
 8 Detroit Red Wings           40  -19   -6   -70

Metropolitan Division
 1 Washington Capitals         40   19    2    25
 2 New York Islanders          37   14    1    13
 3 Pittsburgh Penguins         39   13    6    32
 4 Philadelphia Flyers         39   10    2    11
 5 Carolina Hurricanes         39    9    3    22
 6 New York Rangers            38    4    1    -1
 7 Columbus Blue Jackets       39    3    6   -10
 8 New Jersey Devils           38   -6   -1   -37

Western Conference

Central Division
 1 St. Louis Blues             40   18    6    23
 2 Colorado Avalanche          39   11    0    28
 3 Dallas Stars                40    8    3     6
 4 Winnipeg Jets               39    6   -1     2
 5 Nashville Predators         38    4    1     5
 6 Minnesota Wild              40    3    1    -8
 7 Chicago Blackhawks          40    0    0   -16

Pacific Division
 1 Vancouver Canucks           40    6    2    13
 2 Vegas Golden Knights        42    6    3     6
 3 Arizona Coyotes             41    5   -2     8
 4 Calgary Flames              41    4    1   -11
 5 Edmonton Oilers             41    3   -5   -11
 6 Anaheim Ducks               39   -2   -1   -18
 7 San Jose Sharks             40   -3   -5   -28
 8 Los Angeles Kings           41   -5    2   -25

==================================================
NHL Wildcard Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins               40   17    1    34
 2 Tampa Bay Lightning         37    7    3    15
 3 Florida Panthers            38    7    2     8

Metropolitan Division
 1 Washington Capitals         40   19    2    25
 2 New York Islanders          37   14    1    13
 3 Pittsburgh Penguins         39   13    6    32

Western Conference

Central Division
 1 St. Louis Blues             40   18    6    23
 2 Colorado Avalanche          39   11    0    28
 3 Dallas Stars                40    8    3     6

Pacific Division
 1 Vancouver Canucks           40    6    2    13
 2 Vegas Golden Knights        42    6    3     6
 3 Arizona Coyotes             41    5   -2     8

Eastern Conference Wildcards
 7 Philadelphia Flyers         39   10    2    11
 8 Carolina Hurricanes         39    9    3    22
 9 Toronto Maple Leafs         40    7    7    11
10 New York Rangers            38    4    1    -1
11 Montréal Canadiens          39    3    2     2
12 Columbus Blue Jackets       39    3    6   -10
13 Buffalo Sabres              40    1   -1    -8
14 Ottawa Senators             40   -3    2   -21
15 New Jersey Devils           38   -6   -1   -37
16 Detroit Red Wings           40  -19   -6   -70

Western Conference Wildcards
 7 Winnipeg Jets               39    6   -1     2
 8 Nashville Predators         38    4    1     5
 9 Calgary Flames              41    4    1   -11
10 Minnesota Wild              40    3    1    -8
11 Edmonton Oilers             41    3   -5   -11
12 Chicago Blackhawks          40    0    0   -16
13 Anaheim Ducks               39   -2   -1   -18
14 San Jose Sharks             40   -3   -5   -28
15 Los Angeles Kings           41   -5    2   -25

==================================================
NHL Conference Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference
 1 Washington Capitals         40   19    2    25
 2 Boston Bruins               40   17    1    34
 3 New York Islanders          37   14    1    13
 4 Pittsburgh Penguins         39   13    6    32
 5 Philadelphia Flyers         39   10    2    11
 6 Carolina Hurricanes         39    9    3    22
 7 Tampa Bay Lightning         37    7    3    15
 8 Florida Panthers            38    7    2     8
 9 Toronto Maple Leafs         40    7    7    11
10 New York Rangers            38    4    1    -1
11 Montréal Canadiens          39    3    2     2
12 Columbus Blue Jackets       39    3    6   -10
13 Buffalo Sabres              40    1   -1    -8
14 Ottawa Senators             40   -3    2   -21
15 New Jersey Devils           38   -6   -1   -37
16 Detroit Red Wings           40  -19   -6   -70

Western Conference
 1 St. Louis Blues             40   18    6    23
 2 Colorado Avalanche          39   11    0    28
 3 Dallas Stars                40    8    3     6
 4 Winnipeg Jets               39    6   -1     2
 5 Vancouver Canucks           40    6    2    13
 6 Vegas Golden Knights        42    6    3     6
 7 Arizona Coyotes             41    5   -2     8
 8 Nashville Predators         38    4    1     5
 9 Calgary Flames              41    4    1   -11
10 Minnesota Wild              40    3    1    -8
11 Edmonton Oilers             41    3   -5   -11
12 Chicago Blackhawks          40    0    0   -16
13 Anaheim Ducks               39   -2   -1   -18
14 San Jose Sharks             40   -3   -5   -28
15 Los Angeles Kings           41   -5    2   -25

==================================================
NHL League Standings
==================================================
	                       GP  +/-  +/10   GD
 1 Washington Capitals         40   19    2    25
 2 St. Louis Blues             40   18    6    23
 3 Boston Bruins               40   17    1    34
 4 New York Islanders          37   14    1    13
 5 Pittsburgh Penguins         39   13    6    32
 6 Colorado Avalanche          39   11    0    28
 7 Philadelphia Flyers         39   10    2    11
 8 Carolina Hurricanes         39    9    3    22
 9 Dallas Stars                40    8    3     6
10 Tampa Bay Lightning         37    7    3    15
11 Florida Panthers            38    7    2     8
12 Toronto Maple Leafs         40    7    7    11
13 Winnipeg Jets               39    6   -1     2
14 Vancouver Canucks           40    6    2    13
15 Vegas Golden Knights        42    6    3     6
16 Arizona Coyotes             41    5   -2     8
17 Nashville Predators         38    4    1     5
18 New York Rangers            38    4    1    -1
19 Calgary Flames              41    4    1   -11
20 Montréal Canadiens          39    3    2     2
21 Columbus Blue Jackets       39    3    6   -10
22 Minnesota Wild              40    3    1    -8
23 Edmonton Oilers             41    3   -5   -11
24 Buffalo Sabres              40    1   -1    -8
25 Chicago Blackhawks          40    0    0   -16
26 Anaheim Ducks               39   -2   -1   -18
27 Ottawa Senators             40   -3    2   -21
28 San Jose Sharks             40   -3   -5   -28
29 Los Angeles Kings           41   -5    2   -25
30 New Jersey Devils           38   -6   -1   -37
31 Detroit Red Wings           40  -19   -6   -70

==================================================
NHL Hot or Not, last 10
==================================================
	                       GP  +/-  +/10   GD
 1 Toronto Maple Leafs         40    7    7    11
 2 Pittsburgh Penguins         39   13    6    32
 3 Columbus Blue Jackets       39    3    6   -10
 4 St. Louis Blues             40   18    6    23
 5 Tampa Bay Lightning         37    7    3    15
 6 Carolina Hurricanes         39    9    3    22
 7 Dallas Stars                40    8    3     6
 8 Vegas Golden Knights        42    6    3     6
 9 Florida Panthers            38    7    2     8
10 Philadelphia Flyers         39   10    2    11
11 Montréal Canadiens          39    3    2     2
12 Washington Capitals         40   19    2    25
13 Vancouver Canucks           40    6    2    13
14 Ottawa Senators             40   -3    2   -21
15 Los Angeles Kings           41   -5    2   -25
16 New York Islanders          37   14    1    13
17 Nashville Predators         38    4    1     5
18 New York Rangers            38    4    1    -1
19 Boston Bruins               40   17    1    34
20 Minnesota Wild              40    3    1    -8
21 Calgary Flames              41    4    1   -11
22 Colorado Avalanche          39   11    0    28
23 Chicago Blackhawks          40    0    0   -16
24 New Jersey Devils           38   -6   -1   -37
25 Winnipeg Jets               39    6   -1     2
26 Anaheim Ducks               39   -2   -1   -18
27 Buffalo Sabres              40    1   -1    -8
28 Arizona Coyotes             41    5   -2     8
29 San Jose Sharks             40   -3   -5   -28
30 Edmonton Oilers             41    3   -5   -11
31 Detroit Red Wings           40  -19   -6   -70

```

## Usage (command line flags)

```bash
$ ./nhl -h

Usage of ./nhl:
  -h prints this help info
  -a prints the author information
  -d prints a description of this utility
  -v prints current ersion
```

### Tech used

* https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record

* https://github.com/thedevsaddam/gojsonq/wiki/Queries

* https://golang.org/pkg/sort/