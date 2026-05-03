package sql

const (
	INSERT_PLAYER = `
		INSERT INTO tip_players (api_id, name, short_name, country_code, country_name, age, plays, turned_pro)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	SELECT_PLAYER_BY_ID = `
		SELECT id, api_id, name, short_name, country_code, country_name, age, plays, turned_pro
		FROM tip_players
		WHERE id = ?
	`

	UPDATE_PLAYER = `
		UPDATE tip_players
		SET api_id = ?, name = ?, short_name = ?, country_code = ?, country_name = ?, age = ?, plays = ?, turned_pro = ?
		WHERE id = ?
	`

	DELETE_PLAYER = `
		DELETE FROM tip_players
		WHERE id = ?
	`
)
