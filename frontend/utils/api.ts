// 1. Define the URL


// OLD:
// const API_URL = "http://localhost:8080";

// NEW:
const API_URL = "https://spark-todo.onrender.com";

export interface Todo {
    ID: number;
    title: string;
    completed: boolean;
}

// 2. Function to Get Tasks
// frontend/utils/api.ts

export const getTodos = async (): Promise<Todo[]> => {
    try {
        const response = await fetch(`${API_URL}/todos`);
        
        if (!response.ok) {
        console.error("Backend refused request:", response.status);
        return [];
        }

        const data = await response.json();

        return data.todos || []; 
    } catch (error) {
        console.error("Network error:", error);
        return [];
    }
};

// 3. Function to Create Task
export const createTodo = async (title: string): Promise<Todo> => {
    const response = await fetch(`${API_URL}/todos`, {
        method: "POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({ title }),
    });
    const data = await response.json();
    return data.todo;
};