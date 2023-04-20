INSERT INTO olympics.t_athletes(name, country, sport age, year, date, gold, silver, bronze)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT(name, year, sport) DO
UPDATE SET gold = EXCLUDED.gold + $7, silver = EXCLUDED.silver + $8, bronze = EXCLUDED.silver + $9;