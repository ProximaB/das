-- This script should only run in the very first installation, or rebuild the entire database from scratch.
-- Usage:
--    $ cd dasdb/
--    $ psql -d das -f build.sql -U dasdev
--
-- configure database settings
\! echo "Create DAS Schema"
SET TIME ZONE 'UTC';

-- drop existing schemas
DROP SCHEMA IF EXISTS APPLICATION, COLLEGIATE, DAS, RATING, TMP, USADANCE CASCADE;

-- create all the schemas
CREATE SCHEMA IF NOT EXISTS APPLICATION;
CREATE SCHEMA IF NOT EXISTS COLLEGIATE;
CREATE SCHEMA IF NOT EXISTS DAS;
CREATE SCHEMA IF NOT EXISTS RATING;
CREATE SCHEMA IF NOT EXISTS TMP;
CREATE SCHEMA IF NOT EXISTS USADANCE;

-- create das tables, ordered by dependency, not alphabetically
\i 'tables/das/gender.sql'
\i 'tables/das/account_status.sql'
\i 'tables/das/account_type.sql'
\i 'tables/das/account.sql'
\i 'tables/das/account_role.sql'
\i 'tables/das/account_role_application.sql'
\i 'tables/das/user_preference.sql'
\i 'tables/das/profile.sql'
\i 'tables/das/account_security.sql'

-- create organizer provision tables, under das schema
\i 'tables/das/organizer_provision.sql'
\i 'tables/das/organizer_provision_history.sql'

-- create tables under application schema
\i 'tables/application/version.sql'
\i 'tables/application/announcement_category.sql'
\i 'tables/application/announcement.sql'

-- reference data section under das schema
\i 'tables/das/reference.sql'
\i 'tables/das/city.sql'
\i 'tables/das/school.sql'
\i 'tables/das/studio.sql'
\i 'tables/das/federation.sql'
\i 'tables/das/division.sql'
\i 'tables/das/age.sql'
\i 'tables/das/proficiency.sql'
\i 'tables/das/style.sql'
\i 'tables/das/dance.sql'

-- partnership section under das schema
\i 'tables/das/partnership_request_blacklist_reason.sql'
\i 'tables/das/partnership_request_status.sql'
\i 'tables/das/partnership_request_blacklist.sql'
\i 'tables/das/partnership_role.sql'
\i 'tables/das/partnership_request.sql'
\i 'tables/das/partnership.sql'

-- competition section under das schema
\i 'tables/das/competition_status.sql'
\i 'tables/das/competition.sql'
\i 'tables/das/competition_official.sql'

-- competition product section
\i 'tables/das/competition_product_category.sql'
\i 'tables/das/competition_product_status.sql'
\i 'tables/das/competition_product.sql'
\i 'tables/das/competition_product_order.sql'

-- event section
\i 'tables/das/event_status.sql'
\i 'tables/das/event_category.sql'
\i 'tables/das/event.sql'
\i 'tables/das/event_dances.sql'
\i 'tables/das/event_template.sql'

-- round section
\i 'tables/das/round.sql'

-- entry section
\i 'tables/das/competition_entry.sql'
\i 'tables/das/competition_entry_tba.sql'
\i 'tables/das/competition_representation.sql'
\i 'tables/das/event_entry.sql'
\i 'tables/das/round_entry.sql'

-- round section
\i 'tables/das/event_rounds.sql'

-- scorehsheet section
\i 'tables/das/scoresheet.sql'

-- rank section
\i 'tables/rating/rank_competitive_ballroom.sql'