'use client';

import { useState, useEffect } from 'react';
import AddTodoForm from '@/components/AddTodoForm';
import TodoItem from '@/components/TodoItem';
import EditTodoModal from '@/components/EditTodoModal';

export default function DashboardPage() {
  const [todos, setTodos] = useState([]);
  const [editingTodo, setEditingTodo] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    const stored = localStorage.getItem('todos');
    if (stored) setTodos(JSON.parse(stored));
  }, []);

  useEffect(() => {
    localStorage.setItem('todos', JSON.stringify(todos));
  }, [todos]);

  const addTodo = (text) => {
    const newTodo = { id: Date.now(), text, completed: false };
    setTodos([...todos, newTodo]);
  };

  const toggleComplete = (id) => {
    setTodos(todos.map(todo => todo.id === id ? { ...todo, completed: !todo.completed } : todo));
  };

  const deleteTodo = (id) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };

  const openEditModal = (todo) => {
    setEditingTodo(todo);
    setIsModalOpen(true);
  };

  const closeEditModal = () => {
    setIsModalOpen(false);
    setEditingTodo(null);
  };

  const updateTodo = (id, newText) => {
    setTodos(todos.map(todo => todo.id === id ? { ...todo, text: newText } : todo));
    closeEditModal();
  };

  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 py-8 px-4">
      <div className="max-w-2xl mx-auto bg-white dark:bg-gray-800 rounded-xl shadow-md p-6">
        <h1 className="text-3xl font-bold text-center text-gray-800 dark:text-white mb-4">Maless Dashboard</h1>
        <p className="text-center text-gray-500 dark:text-gray-400 mb-6">Your personal todo list to stay productive 🚀</p>
        
        <AddTodoForm addTodo={addTodo} />

        <div className="mt-6 space-y-3">
          {todos.length === 0 ? (
            <p className="text-center text-gray-400">No tasks yet. Add one above!</p>
          ) : (
            todos.map(todo => (
              <TodoItem
                key={todo.id}
                todo={todo}
                onToggle={toggleComplete}
                onDelete={deleteTodo}
                onEdit={openEditModal}
              />
            ))
          )}
        </div>
      </div>

      {/* Modal */}
      {isModalOpen && editingTodo && (
        <EditTodoModal 
          todo={editingTodo} 
          onClose={closeEditModal} 
          onUpdate={updateTodo} 
        />
      )}
    </main>
  );
}
