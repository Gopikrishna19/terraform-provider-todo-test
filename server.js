const express = require('express');
const bodyParser = require("body-parser");

const app = express();
const todos = [];
let id = 0;

const findTodo = (req, res, next) => {
    const uId = parseInt(req.params.id);
    const index = todos.findIndex(t => t.id === uId);

    if (index === -1) {
        res.statusCode = 404;
        res.send({})
    } else {
        const todo = {
            ...todos[index]
        };

        req.target = {
            todo,
            index
        };

        next();
    }
}

app.use(bodyParser.json());

app.get('/todo', (req, res) => {
    res.send(todos);
});

app.post('/todo', (req, res) => {
    id += 1;
    const todo = {
        id,
        completed: false,
        task: req.body.task
    };

    todos.push(todo)

    res.send({id});
});

app.put('/todo/:id', findTodo, (req, res) => {
    const {index, todo} = req.target;

    todo.completed = req.body.completed;
    todos.splice(index, 1, todo);

    res.send(todo);
});

app.delete('/todo/:id', findTodo, (req, res) => {
    const {index} = req.target;

    todos.splice(index, 1);

    res.send({});
});

app.listen(8080);
