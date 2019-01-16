select
  ecb.id as event_id,
  comp.name as competition,
  comp.datetime_start as start,
  ec.name as event_category,
  f.name as federation,
  d.name as division,
  p.name as proficiency,
  s.name as style,
  da.name as dance
from
  das.event_competitive_ballroom ecb
join das.event e on ecb.event_id = e.id
join das.competition comp on e.competition_id = comp.id
join das.federation f on ecb.federation_id = f.id
join das.division d on ecb.division_id = d.id
join das.proficiency p on ecb.proficiency_id = p.id
join das.style s on ecb.style_id = s.id
join das.event_competitive_ballroom_dances ecbd on ecbd.competitive_ballroom_event_id = ecb.id
join das.dance da on ecbd.dance_id = da.id
join das.event_category ec on e.event_category_id = ec.id;