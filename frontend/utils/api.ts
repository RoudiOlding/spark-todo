// 1. Define the URL

const API_URL = "http://localhost:8080";

export interface Todo {
    ID: number;
    title: string;
    completed: boolean;
}

// 2. Function to Get Tasks
export const getTodos = async (): Promise<Todo[]> => {
    const response = await fetch(`${API_URL}/todos`);
    const data = await response.json();
    return data.todos;
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