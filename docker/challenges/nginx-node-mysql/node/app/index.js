const express = require('express');
const mysql = require('mysql2');

const app = express();
const port = 3000;

const connection = mysql.createConnection({
    host: 'db',
    user: 'root',
    password: 'root',
    database: 'nodedb'
});

// Create a table called people with two columns: id and name, where id is the PK int and name is a varchar
connection.connect((error) => {
    if (error) {
        console.error('Error connecting to database:', error);
        return;
    }

    const createTableQuery = `CREATE TABLE IF NOT EXISTS people (
        id int not null auto_increment primary key,
        name varchar(255)
    )`;

    connection.query(createTableQuery, (error) => {
        if (error) {
            console.error('Error creating table:', error);
            return;
        }

        console.log('Table created/updated successfully');
    });
});

app.get('/', (req, res) => {
    const insertQuery = `INSERT INTO people (name) VALUES ('John Doe')`;

    connection.query(insertQuery, (error, results) => {
        if (error) {
            console.error('Error inserting data:', error);
            res.status(500).send('Error inserting data');
            return;
        }
        console.log('Data inserted successfully');

        const selectQuery = 'SELECT * FROM people';

        connection.query(selectQuery, (error, results) => {
            if (error) {
                console.error('Error querying data:', error);
                res.status(500).send('Error querying data');
                return;
            }
            console.log('Data selected successfully');
            res.send(
                `<h1>Full Cycle Rocks!</h1>
                ${results.map(item => {
                    return `<p>- ${item.name}</p>`;
                }).join('')
                }`
            );
        });
    });
});

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});