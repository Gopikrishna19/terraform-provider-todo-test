provider "todo-test" {
  host = "http://localhost"
  port = "8080"
}

resource "todo" "todo1" {
  task = "Buy milk"
}
  
