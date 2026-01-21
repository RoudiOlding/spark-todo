"use client";

import { useState, useEffect } from "react";
import { getTodos, createTodo, Todo, updateTodo, deleteTodo } from "../utils/api";

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

  // 4. Handle "Completed Task"
  const handleToggle = async (id: number, currentStatus: boolean) => {
    const updatedTodos = todos.map((t) => 
      t.ID === id ? { ...t, completed: !currentStatus } : t
    );
    setTodos(updatedTodos);

    // 2. Call API
    await updateTodo(id, !currentStatus);
  };

  const handleDelete = async (id: number) => {
    setTodos(todos.filter((t) => t.ID !== id));

    await deleteTodo(id);
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
          // 1. Just ONE main container for the card
          <div
            key={todo.ID}
            className="flex items-center p-4 bg-gray-800 rounded border border-gray-700 shadow-sm transition-all hover:bg-gray-750"
          >
            {/* 2. The Button (Circle) */}
            <button
              onClick={() => handleToggle(todo.ID, todo.completed)}
              className={`w-6 h-6 rounded-full mr-4 border-2 flex items-center justify-center transition-all ${
                todo.completed 
                  ? "bg-green-500 border-green-500" 
                  : "border-gray-500 hover:border-blue-400"
              }`}
            >
              {todo.completed && "✓"}
            </button>
            
            {/* 3. The Title */}
            <span className={todo.completed ? "line-through text-gray-500" : "text-gray-100"}>
              {todo.title}
            </span>

            <button 
              onClick={() => handleDelete(todo.ID)}
              className="ml-4 text-gray-500 hover:text-red-500 transition-colors p-2"
              aria-label="Delete task"
            >
              🗑️
            </button>
          </div>
        ))}
        
        {todos.length === 0 && (
          <p className="text-center text-gray-500 mt-10">No tasks yet. Start building!</p>
        )}
      </div>

    </main>
  );
}