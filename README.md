# economic-os

Aggregates economic data from BLS, BEA, and other sources.
Displays data in a dashboard, showing an aggregated state of the economy.
Schedule collects economic data release dates, available as ics calendar to subscribe to.

Current todo:
- release calendar, collect dates from sources to know when data will release to. collect in db, then work on making subscribe-able ics
- once release calendar properly added, then ensure pulling data from each source works
- after data connections established, work on architecture for finding release date -> pulling release data -> saving data locally so don't need to pull off of api constantly
- after that, work on how to display everything on frontend, dashboard for data and releases, ability to create blog posts
