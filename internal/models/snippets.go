package models

import (
	"database/sql"
	"errors"
	"time"
)

//Define a Snippet type to hold data for an individual snippet

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

//Define  SnippetModel type which wraps a sql.DB connection pool

type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet int the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	//The Exec() method on execution of statments on the connection pool
	//The first param is the statment followed by the params the
	//returned type is sq.Result whcih contains info on what happened when teh statment was exected
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	//use the LastInsert ID methos to get the id of our record on the table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

// This will return a new snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	//USE THE QUERYROW() METHOD ON CONNECTION POOL TO EXCUTRE OUR SQL STATMENT
	row := m.DB.QueryRow(stmt, id)

	//Initialize a pointer to a new zeroed snippet struct
	s := &Snippet{}

	//using row.Scan() to copy the values from each field in sq.Row to
	//the corresponfing field in Snippet struct.The arguments are pointer to the data

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

//This will return 10 most recently create snippets

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	//use the Query() nethod on the  connection pool to execute our SQL statment.
	//This returns a sql.Rows resultset contatining the result od our query

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.CLose() to ensure the resultset is closed before Latest() method returns.
	//It should come after we check an erro.Because they will get a panic when trying to close a nil resultset.
	defer rows.Close()

	//Initialize an empty slice to hold the snippet structs
	snippets := []*Snippet{}
	//use rows.Next to iterate through the rows in resultSet.this prepares the row to prepared
	//the  first and each subsequent row to be acted on by rows.Scan() methos.If iteration over all the rows completes then the resultset automatically closes and frees up the the underlying database connection

	for rows.Next() {
		s := &Snippet{}
		//use the rows.Scan() methos to copy values
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	//we call rows.Err() to check if iteration causes an error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
