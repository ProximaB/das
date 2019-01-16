select a.id, a.first_name, a.last_name, t.name from das.account a 
join das.account_role r on a.id = r.account_id 
join das.account_type t on t.id = r.account_type_id
order by a.id, r.account_type_id;