import express from 'express';

const app = express();

const PORT = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
    res.status(200).send('hello !');
});

app.get('/getit', (req, res) => {
    res.status(200).json({ message: 'hello from `/get`' });
});

app.post('/post', (req, res) => {
    const myJSON = req.body;
    res.status(200).json({ message: myJSON });
});

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => console.log(`lsitening to: http://localhost:${PORT}`));
