"use client"

import { motion } from "framer-motion"
import { Trash, Edit } from "lucide-react"

export default function TodoItem({ todo, onDelete, onToggle, onEdit }) {
  return (
    <motion.div
      layout
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, height: 0, marginBottom: 0 }}
      transition={{ duration: 0.2 }}
      className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:shadow transition-all"
    >
      <div className="flex items-center gap-3">
        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() => onToggle(todo.id)}
          className="w-5 h-5 rounded-full border-2 border-gray-300 dark:border-gray-600 focus:ring-2 focus:ring-blue-500 text-blue-600 cursor-pointer"
        />
        <span
          className={`${
            todo.completed ? "line-through text-gray-400 dark:text-gray-500" : "text-gray-800 dark:text-gray-200"
          }`}
        >
          {todo.text}
        </span>
      </div>
      <div className="flex gap-2">
        <button
          onClick={() => onEdit(todo)}
          className="p-1 text-gray-500 hover:text-blue-500 dark:text-gray-400 dark:hover:text-blue-400 transition-colors"
          aria-label="Edit task"
        >
          <Edit size={18} />
        </button>
        <button
          onClick={() => onDelete(todo.id)}
          className="p-1 text-gray-500 hover:text-red-500 dark:text-gray-400 dark:hover:text-red-400 transition-colors"
          aria-label="Delete task"
        >
          <Trash size={18} />
        </button>
      </div>
    </motion.div>
  )
}

