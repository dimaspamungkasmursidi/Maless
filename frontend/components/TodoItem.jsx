'use client';

export default function TodoItem({ todo, onToggle, onDelete, onEdit }) {
  return (
    <div className="flex items-center justify-between p-4 bg-white dark:bg-gray-700 rounded-lg shadow-md">
      <div className="flex items-center">
        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() => onToggle(todo.id)}
          className="mr-3"
        />
        <span
          className={`text-lg ${todo.completed ? 'line-through text-gray-400' : 'text-gray-800 dark:text-white'}`}
        >
          {todo.text}
        </span>
      </div>
      <div className="flex items-center space-x-2">
        <button
          onClick={() => onEdit(todo)}
          className="text-blue-500 hover:text-blue-700"
        >
          Edit
        </button>
        <button
          onClick={() => onDelete(todo.id)}
          className="text-red-500 hover:text-red-700"
        >
          Delete
        </button>
      </div>
    </div>
  );
}
