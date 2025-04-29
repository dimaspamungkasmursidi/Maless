'use client';

import { useState } from 'react';

export default function EditTodoModal({ todo, onClose, onUpdate }) {
  const [newText, setNewText] = useState(todo.text);

  const handleUpdate = (e) => {
    e.preventDefault();
    if (newText.trim()) {
      onUpdate(todo.id, newText.trim());
      onClose();
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
      <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-semibold text-gray-800 dark:text-white mb-4">Edit Todo</h2>
        
        <form onSubmit={handleUpdate}>
          <input
            type="text"
            value={newText}
            onChange={(e) => setNewText(e.target.value)}
            className="w-full px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-400 dark:bg-gray-700 dark:text-white mb-4"
            placeholder="Edit your todo"
          />
          <div className="flex justify-between">
            <button
              type="button"
              onClick={onClose}
              className="bg-gray-400 text-white px-4 py-2 rounded-md hover:bg-gray-500"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700"
            >
              Update
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
