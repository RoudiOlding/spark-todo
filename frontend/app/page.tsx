"use client";

import { useState, useEffect } from "react";
import { getTodos, createTodo, Todo } from "../utils/api";

export default function Home() {
  // 1. Define State (The memory of the screen)
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTitle, setNewTitle] = useState("");

  // 2. Load Tasks on Start 
  useEffect(() => {
    getTodos().then((data) => {
      setTodos(data);
    });
  }, []);

  // 3. Handle "Add Task"
  const handleAdd = async () => {
    if (!newTitle) return;
    
    const created = await createTodo(newTitle);
    
    setTodos([...todos, created]);
    setNewTitle("");
  };

  return (
    <main className="min-h-screen bg-gray-900 text-white flex flex-col items-center p-10">
      <h1 className="text-4xl font-bold mb-8 text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-500">
        Sparks Studio To-Do
      </h1>

      {/* Input Section */}
      <div className="flex gap-2 mb-8 w-full max-w-md">
        <input
          type="text"
          value={newTitle}
          onChange={(e) => setNewTitle(e.target.value)}
          placeholder="What needs to be done?"
          className="flex-1 p-3 rounded bg-gray-800 border border-gray-700 focus:outline-none focus:border-blue-500"
          onKeyDown={(e) => e.key === "Enter" && handleAdd()}
        />
        <button
          onClick={handleAdd}
          className="bg-blue-600 hover:bg-blue-500 px-6 py-3 rounded font-bold transition-all"
        >
          Add
        </button>
      </div>

      {/* List Section */}
      <div className="w-full max-w-md space-y-3">
        {todos.map((todo) => (
          <div
            key={todo.ID}
            className="flex items-center p-4 bg-gray-800 rounded border border-gray-700 shadow-sm"
          >
            {/* Circle Icon */}
            <div className={`w-4 h-4 rounded-full mr-4 ${todo.completed ? "bg-green-500" : "bg-gray-600"}`}></div>
            
            <span className={todo.completed ? "line-through text-gray-500" : "text-gray-100"}>
              {todo.title}
            </span>
          </div>
        ))}
        
        {todos.length === 0 && (
          <p className="text-center text-gray-500 mt-10">No tasks yet. Start building!</p>
        )}
      </div>
    </main>
  );
}