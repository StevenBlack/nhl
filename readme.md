# NHL team plus/minus standings.

I'm unmoved by traditional National Hockey League (NHL) reporting
which tallies standings by points.

Given the disparity, league wide, in games played, total points
is always misleading.

The only metric that matters is, wins minus losses.

To make the playoffs, generally 96 points gets you there.
Given an 82-game regular season, 96 points is +14 — a differential of +14, wins minus losses.

To run this:

* `./refresh` fetches stats and refreshes local data.
* `./nhl` lists various facets of NHL standings, using local data.


### Tech used

* https://statsapi.web.nhl.com/api/v1/standings?expand=standings.record

* https://github.com/thedevsaddam/gojsonq/wiki/Queries

* https://golang.org/pkg/sort/