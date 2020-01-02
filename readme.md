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

============================================
NHL Division Standings
============================================
	                  +/-  +/10  GP   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins           17    2   41   33
 2 Tampa Bay Lightning      8    5   38   17
 3 Toronto Maple Leafs      8    7   41   14
 4 Florida Panthers         6    0   39    5
 5 Montréal Canadiens       2    0   40    0
 6 Buffalo Sabres           0   -3   41  -10
 7 Ottawa Senators         -3    2   40  -21
 8 Detroit Red Wings      -18   -4   41  -68

Metropolitan Division
 1 Washington Capitals     18    0   41   24
 2 New York Islanders      15    3   38   14
 3 Pittsburgh Penguins     13    6   39   32
 4 Carolina Hurricanes     10    3   40   24
 5 Philadelphia Flyers      9    0   40    9
 6 Columbus Blue Jackets    4    6   40   -7
 7 New York Rangers         3   -1   39   -3
 8 New Jersey Devils       -5    1   39  -36

Western Conference

Central Division
 1 St. Louis Blues         17    6   41   21
 2 Colorado Avalanche      10   -1   40   25
 3 Dallas Stars             9    3   41    8
 4 Winnipeg Jets            7   -1   40    5
 5 Nashville Predators      3   -1   39    3
 6 Minnesota Wild           2    0   41  -11
 7 Chicago Blackhawks       1    2   41  -14

Pacific Division
 1 Vegas Golden Knights     7    3   43    9
 2 Vancouver Canucks        6    2   40   13
 3 Arizona Coyotes          6   -2   42   10
 4 Edmonton Oilers          4   -4   42   -9
 5 Calgary Flames           3   -1   42  -13
 6 Anaheim Ducks           -3   -1   40  -21
 7 San Jose Sharks         -4   -5   41  -30
 8 Los Angeles Kings       -4    2   42  -23

============================================
NHL Wildcard Standings
============================================
	                  +/-  +/10  GP   GD

Eastern Conference

Atlantic Division
 1 Boston Bruins           17    2   41   33
 2 Tampa Bay Lightning      8    5   38   17
 3 Toronto Maple Leafs      8    7   41   14

Metropolitan Division
 1 Washington Capitals     18    0   41   24
 2 New York Islanders      15    3   38   14
 3 Pittsburgh Penguins     13    6   39   32

Western Conference

Central Division
 1 St. Louis Blues         17    6   41   21
 2 Colorado Avalanche      10   -1   40   25
 3 Dallas Stars             9    3   41    8

Pacific Division
 1 Vegas Golden Knights     7    3   43    9
 2 Vancouver Canucks        6    2   40   13
 3 Arizona Coyotes          6   -2   42   10

Eastern Conference Wildcards
 7 Carolina Hurricanes     10    3   40   24
 8 Philadelphia Flyers      9    0   40    9
 9 Florida Panthers         6    0   39    5
10 Columbus Blue Jackets    4    6   40   -7
11 New York Rangers         3   -1   39   -3
12 Montréal Canadiens       2    0   40    0
13 Buffalo Sabres           0   -3   41  -10
14 Ottawa Senators         -3    2   40  -21
15 New Jersey Devils       -5    1   39  -36
16 Detroit Red Wings      -18   -4   41  -68

Western Conference Wildcards
 7 Winnipeg Jets            7   -1   40    5
 8 Edmonton Oilers          4   -4   42   -9
 9 Nashville Predators      3   -1   39    3
10 Calgary Flames           3   -1   42  -13
11 Minnesota Wild           2    0   41  -11
12 Chicago Blackhawks       1    2   41  -14
13 Anaheim Ducks           -3   -1   40  -21
14 San Jose Sharks         -4   -5   41  -30
15 Los Angeles Kings       -4    2   42  -23

============================================
NHL Conference Standings
============================================
	                  +/-  +/10  GP   GD

Eastern Conference
 1 Washington Capitals     18    0   41   24
 2 Boston Bruins           17    2   41   33
 3 New York Islanders      15    3   38   14
 4 Pittsburgh Penguins     13    6   39   32
 5 Carolina Hurricanes     10    3   40   24
 6 Philadelphia Flyers      9    0   40    9
 7 Tampa Bay Lightning      8    5   38   17
 8 Toronto Maple Leafs      8    7   41   14
 9 Florida Panthers         6    0   39    5
10 Columbus Blue Jackets    4    6   40   -7
11 New York Rangers         3   -1   39   -3
12 Montréal Canadiens       2    0   40    0
13 Buffalo Sabres           0   -3   41  -10
14 Ottawa Senators         -3    2   40  -21
15 New Jersey Devils       -5    1   39  -36
16 Detroit Red Wings      -18   -4   41  -68

Western Conference
 1 St. Louis Blues         17    6   41   21
 2 Colorado Avalanche      10   -1   40   25
 3 Dallas Stars             9    3   41    8
 4 Winnipeg Jets            7   -1   40    5
 5 Vegas Golden Knights     7    3   43    9
 6 Vancouver Canucks        6    2   40   13
 7 Arizona Coyotes          6   -2   42   10
 8 Edmonton Oilers          4   -4   42   -9
 9 Nashville Predators      3   -1   39    3
10 Calgary Flames           3   -1   42  -13
11 Minnesota Wild           2    0   41  -11
12 Chicago Blackhawks       1    2   41  -14
13 Anaheim Ducks           -3   -1   40  -21
14 San Jose Sharks         -4   -5   41  -30
15 Los Angeles Kings       -4    2   42  -23

============================================
NHL League Standings
============================================
	                  +/-  +/10  GP   GD
 1 Washington Capitals     18    0   41   24
 2 Boston Bruins           17    2   41   33
 3 St. Louis Blues         17    6   41   21
 4 New York Islanders      15    3   38   14
 5 Pittsburgh Penguins     13    6   39   32
 6 Colorado Avalanche      10   -1   40   25
 7 Carolina Hurricanes     10    3   40   24
 8 Philadelphia Flyers      9    0   40    9
 9 Dallas Stars             9    3   41    8
10 Tampa Bay Lightning      8    5   38   17
11 Toronto Maple Leafs      8    7   41   14
12 Winnipeg Jets            7   -1   40    5
13 Vegas Golden Knights     7    3   43    9
14 Florida Panthers         6    0   39    5
15 Vancouver Canucks        6    2   40   13
16 Arizona Coyotes          6   -2   42   10
17 Columbus Blue Jackets    4    6   40   -7
18 Edmonton Oilers          4   -4   42   -9
19 Nashville Predators      3   -1   39    3
20 New York Rangers         3   -1   39   -3
21 Calgary Flames           3   -1   42  -13
22 Montréal Canadiens       2    0   40    0
23 Minnesota Wild           2    0   41  -11
24 Chicago Blackhawks       1    2   41  -14
25 Buffalo Sabres           0   -3   41  -10
26 Anaheim Ducks           -3   -1   40  -21
27 Ottawa Senators         -3    2   40  -21
28 San Jose Sharks         -4   -5   41  -30
29 Los Angeles Kings       -4    2   42  -23
30 New Jersey Devils       -5    1   39  -36
31 Detroit Red Wings      -18   -4   41  -68

============================================
NHL Hot or Not, last 10
============================================
	                  +/-  +/10  GP   GD
 1 Toronto Maple Leafs      8    7   41   14
 2 Pittsburgh Penguins     13    6   39   32
 3 Columbus Blue Jackets    4    6   40   -7
 4 St. Louis Blues         17    6   41   21
 5 Tampa Bay Lightning      8    5   38   17
 6 New York Islanders      15    3   38   14
 7 Carolina Hurricanes     10    3   40   24
 8 Dallas Stars             9    3   41    8
 9 Vegas Golden Knights     7    3   43    9
10 Vancouver Canucks        6    2   40   13
11 Ottawa Senators         -3    2   40  -21
12 Boston Bruins           17    2   41   33
13 Chicago Blackhawks       1    2   41  -14
14 Los Angeles Kings       -4    2   42  -23
15 New Jersey Devils       -5    1   39  -36
16 Florida Panthers         6    0   39    5
17 Philadelphia Flyers      9    0   40    9
18 Montréal Canadiens       2    0   40    0
19 Washington Capitals     18    0   41   24
20 Minnesota Wild           2    0   41  -11
21 Nashville Predators      3   -1   39    3
22 New York Rangers         3   -1   39   -3
23 Colorado Avalanche      10   -1   40   25
24 Winnipeg Jets            7   -1   40    5
25 Anaheim Ducks           -3   -1   40  -21
26 Calgary Flames           3   -1   42  -13
27 Arizona Coyotes          6   -2   42   10
28 Buffalo Sabres           0   -3   41  -10
29 Detroit Red Wings      -18   -4   41  -68
30 Edmonton Oilers          4   -4   42   -9
31 San Jose Sharks         -4   -5   41  -30

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