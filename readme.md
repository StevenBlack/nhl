# NHL wins minus losses plus/minus standings

This is a personal side-project, written in [Go](https://golang.org/).

I'm unmoved by traditional National Hockey League (NHL) reporting
which [tallies standings by points](https://www.nhl.com/standings/).

Given the disparity, league wide, in games played throughout the season, 
total points ranking is always misleading.

The only metric that matters is, wins minus losses.

To make the playoffs, generally a team needs 96 points to make it.
Given an 82-game regular season, 96 points is +14 — a differential of +14, wins minus losses.

This is what this repo gives you: plaintext plus/minus NHL standings in the command window.

To run this:

* `./refresh` fetches stats and refreshes local data.
* `./nhl` lists various facets of NHL standings, using local data.


## Sample output

Refresh local data:

```
$ ./refresh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 68720    0 68720    0     0   309k      0 --:--:-- --:--:-- --:--:--  309k
```

Show NHL plus/minus team standings in various ways.

* `GP`: games played.
* `+/-`: plus/minus, season so far.
* `+/10`: plus/minus, last 10 games played.
* `GD`: team goal differential, season so far.

```
$ ./nhl

==================================================
NHL Division Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins               39   16   -1    33
 2 Toronto Maple Leafs         40    7    7    11
 3 Tampa Bay Lightning         36    6    3    14
 4 Florida Panthers            37    6    2     7
 5 Montréal Canadiens          38    4    2     3
 6 Buffalo Sabres              39    2    0    -7
 7 Ottawa Senators             38   -2    3   -17
 8 Detroit Red Wings           39  -18   -6   -69

Metropolitan Division
 1 Washington Capitals         40   19    2    25
 2 New York Islanders          36   13    1    11
 3 Pittsburgh Penguins         38   12    6    29
 4 Philadelphia Flyers         38    9    0    10
 5 Carolina Hurricanes         39    9    3    22
 6 New York Rangers            38    4    1    -1
 7 Columbus Blue Jackets       38    3    5    -9
 8 New Jersey Devils           37   -7   -2   -38

Western Conference

Central Division
 1 St. Louis Blues             39   17    4    20
 2 Colorado Avalanche          39   11    0    28
 3 Winnipeg Jets               38    7    0     5
 4 Dallas Stars                39    7    3     4
 5 Nashville Predators         38    4    1     5
 6 Minnesota Wild              39    4    1    -6
 7 Chicago Blackhawks          39   -1   -1   -17

Pacific Division
 1 Arizona Coyotes             40    6   -2    10
 2 Vegas Golden Knights        42    6    3     6
 3 Vancouver Canucks           39    5    2    10
 4 Calgary Flames              40    5    3    -8
 5 Edmonton Oilers             41    3   -5   -11
 6 Anaheim Ducks               38   -2   -2   -17
 7 San Jose Sharks             40   -3   -5   -28
 8 Los Angeles Kings           41   -5    2   -25

==================================================
NHL Wildcard Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins               39   16   -1    33
 2 Toronto Maple Leafs         40    7    7    11
 3 Tampa Bay Lightning         36    6    3    14

Metropolitan Division
 1 Washington Capitals         40   19    2    25
 2 New York Islanders          36   13    1    11
 3 Pittsburgh Penguins         38   12    6    29

Western Conference

Central Division
 1 St. Louis Blues             39   17    4    20
 2 Colorado Avalanche          39   11    0    28
 3 Winnipeg Jets               38    7    0     5

Pacific Division
 1 Arizona Coyotes             40    6   -2    10
 2 Vegas Golden Knights        42    6    3     6
 3 Vancouver Canucks           39    5    2    10

Eastern Conference Wildcards
 7 Philadelphia Flyers         38    9    0    10
 8 Carolina Hurricanes         39    9    3    22
 9 Florida Panthers            37    6    2     7
10 Montréal Canadiens          38    4    2     3
11 New York Rangers            38    4    1    -1
12 Columbus Blue Jackets       38    3    5    -9
13 Buffalo Sabres              39    2    0    -7
14 Ottawa Senators             38   -2    3   -17
15 New Jersey Devils           37   -7   -2   -38
16 Detroit Red Wings           39  -18   -6   -69

Western Conference Wildcards
 7 Dallas Stars                39    7    3     4
 8 Calgary Flames              40    5    3    -8
 9 Nashville Predators         38    4    1     5
10 Minnesota Wild              39    4    1    -6
11 Edmonton Oilers             41    3   -5   -11
12 Chicago Blackhawks          39   -1   -1   -17
13 Anaheim Ducks               38   -2   -2   -17
14 San Jose Sharks             40   -3   -5   -28
15 Los Angeles Kings           41   -5    2   -25

==================================================
NHL Conference Standings
==================================================
	                       GP  +/-  +/10   GD

Eastern Conference
 1 Washington Capitals         40   19    2    25
 2 Boston Bruins               39   16   -1    33
 3 New York Islanders          36   13    1    11
 4 Pittsburgh Penguins         38   12    6    29
 5 Philadelphia Flyers         38    9    0    10
 6 Carolina Hurricanes         39    9    3    22
 7 Toronto Maple Leafs         40    7    7    11
 8 Tampa Bay Lightning         36    6    3    14
 9 Florida Panthers            37    6    2     7
10 Montréal Canadiens          38    4    2     3
11 New York Rangers            38    4    1    -1
12 Columbus Blue Jackets       38    3    5    -9
13 Buffalo Sabres              39    2    0    -7
14 Ottawa Senators             38   -2    3   -17
15 New Jersey Devils           37   -7   -2   -38
16 Detroit Red Wings           39  -18   -6   -69

Western Conference
 1 St. Louis Blues             39   17    4    20
 2 Colorado Avalanche          39   11    0    28
 3 Winnipeg Jets               38    7    0     5
 4 Dallas Stars                39    7    3     4
 5 Arizona Coyotes             40    6   -2    10
 6 Vegas Golden Knights        42    6    3     6
 7 Vancouver Canucks           39    5    2    10
 8 Calgary Flames              40    5    3    -8
 9 Nashville Predators         38    4    1     5
10 Minnesota Wild              39    4    1    -6
11 Edmonton Oilers             41    3   -5   -11
12 Chicago Blackhawks          39   -1   -1   -17
13 Anaheim Ducks               38   -2   -2   -17
14 San Jose Sharks             40   -3   -5   -28
15 Los Angeles Kings           41   -5    2   -25

==================================================
NHL League Standings
==================================================
	                       GP  +/-  +/10   GD
 1 Washington Capitals         40   19    2    25
 2 St. Louis Blues             39   17    4    20
 3 Boston Bruins               39   16   -1    33
 4 New York Islanders          36   13    1    11
 5 Pittsburgh Penguins         38   12    6    29
 6 Colorado Avalanche          39   11    0    28
 7 Philadelphia Flyers         38    9    0    10
 8 Carolina Hurricanes         39    9    3    22
 9 Winnipeg Jets               38    7    0     5
10 Dallas Stars                39    7    3     4
11 Toronto Maple Leafs         40    7    7    11
12 Tampa Bay Lightning         36    6    3    14
13 Florida Panthers            37    6    2     7
14 Arizona Coyotes             40    6   -2    10
15 Vegas Golden Knights        42    6    3     6
16 Vancouver Canucks           39    5    2    10
17 Calgary Flames              40    5    3    -8
18 Nashville Predators         38    4    1     5
19 Montréal Canadiens          38    4    2     3
20 New York Rangers            38    4    1    -1
21 Minnesota Wild              39    4    1    -6
22 Columbus Blue Jackets       38    3    5    -9
23 Edmonton Oilers             41    3   -5   -11
24 Buffalo Sabres              39    2    0    -7
25 Chicago Blackhawks          39   -1   -1   -17
26 Anaheim Ducks               38   -2   -2   -17
27 Ottawa Senators             38   -2    3   -17
28 San Jose Sharks             40   -3   -5   -28
29 Los Angeles Kings           41   -5    2   -25
30 New Jersey Devils           37   -7   -2   -38
31 Detroit Red Wings           39  -18   -6   -69

==================================================
NHL Hot or Not, last 10
==================================================
	                       GP  +/-  +/10   GD
 1 Toronto Maple Leafs         40    7    7    11
 2 Pittsburgh Penguins         38   12    6    29
 3 Columbus Blue Jackets       38    3    5    -9
 4 St. Louis Blues             39   17    4    20
 5 Tampa Bay Lightning         36    6    3    14
 6 Ottawa Senators             38   -2    3   -17
 7 Carolina Hurricanes         39    9    3    22
 8 Dallas Stars                39    7    3     4
 9 Calgary Flames              40    5    3    -8
10 Vegas Golden Knights        42    6    3     6
11 Florida Panthers            37    6    2     7
12 Montréal Canadiens          38    4    2     3
13 Vancouver Canucks           39    5    2    10
14 Washington Capitals         40   19    2    25
15 Los Angeles Kings           41   -5    2   -25
16 New York Islanders          36   13    1    11
17 Nashville Predators         38    4    1     5
18 New York Rangers            38    4    1    -1
19 Minnesota Wild              39    4    1    -6
20 Philadelphia Flyers         38    9    0    10
21 Winnipeg Jets               38    7    0     5
22 Colorado Avalanche          39   11    0    28
23 Buffalo Sabres              39    2    0    -7
24 Boston Bruins               39   16   -1    33
25 Chicago Blackhawks          39   -1   -1   -17
26 New Jersey Devils           37   -7   -2   -38
27 Anaheim Ducks               38   -2   -2   -17
28 Arizona Coyotes             40    6   -2    10
29 San Jose Sharks             40   -3   -5   -28
30 Edmonton Oilers             41    3   -5   -11
31 Detroit Red Wings           39  -18   -6   -69
```

### Tech used

* https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record

* https://github.com/thedevsaddam/gojsonq/wiki/Queries

* https://golang.org/pkg/sort/