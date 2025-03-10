"use client"

import { useState, useEffect } from "react"
import { motion, AnimatePresence } from "framer-motion"
import TodoItem from "@/components/todo-item"
import AddTodoForm from "@/components/add-todo-form"
import EditTodoModal from "@/components/edit-todo-modal"

export default function Home() {
  const [todos, setTodos] = useState([])
  const [editingTodo, setEditingTodo] = useState(null)
  const [isModalOpen, setIsModalOpen] = useState(false)

  // Load todos from localStorage on component mount
  useEffect(() => {
    const storedTodos = localStorage.getItem("todos")
    if (storedTodos) {
      setTodos(JSON.parse(storedTodos))
    }
  }, [])

  // Save todos to localStorage whenever todos change
  useEffect(() => {
    localStorage.setItem("todos", JSON.stringify(todos))
  }, [todos])

  const addTodo = (text) => {
    const newTodo = {
      id: Date.now(),
      text,
      completed: false,
    }
    setTodos([...todos, newTodo])
  }

  const deleteTodo = (id) => {
    setTodos(todos.filter((todo) => todo.id !== id))
  }

  const toggleComplete = (id) => {
    setTodos(todos.map((todo) => (todo.id === id ? { ...todo, completed: !todo.completed } : todo)))
  }

  const openEditModal = (todo) => {
    setEditingTodo(todo)
    setIsModalOpen(true)
  }

  const closeEditModal = () => {
    setIsModalOpen(false)
    setEditingTodo(null)
  }

  const updateTodo = (id, newText) => {
    setTodos(todos.map((todo) => (todo.id === id ? { ...todo, text: newText } : todo)))
    closeEditModal()
  }

  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800 py-8 px-4">
      <motion.div
        initial={{ opacity: 0, y: -20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="max-w-2xl mx-auto bg-white dark:bg-gray-800 rounded-xl shadow-md overflow-hidden"
      >
        <div className="p-6">
          <h1 className="text-2xl font-bold text-center text-gray-800 dark:text-white mb-2">Maless</h1>
          <p className="text-center text-gray-600 dark:text-gray-400"><i>'Malas Less'</i> meaning <i>'less laziness'</i>.</p>
          <p className="text-center text-gray-600 dark:text-gray-400 mb-6">is a to-do list app designed to help you stay productive by reducing laziness.</p>
          {/* <p className="text-center text-gray-600 dark:text-gray-400 mb-6">Maless is a to-do list app designed to help you stay productive by reducing laziness. The name 'Maless' comes from 'Malas Less,' meaning 'less laziness.' With an intuitive interface and smart task management features, Maless keeps you organized, motivated, and on track to achieve your goals effortlessly!.</p> */}
          <AddTodoForm addTodo={addTodo} />

          <div className="mt-6 space-y-2">
            <AnimatePresence>
              {todos.length === 0 ? (
                <motion.p
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 0.6 }}
                  exit={{ opacity: 0 }}
                  className="text-center text-gray-500 dark:text-gray-400 py-4"
                >
                  No tasks yet. Add one above!
                </motion.p>
              ) : (
                todos.map((todo) => (
                  <TodoItem
                    key={todo.id}
                    todo={todo}
                    onDelete={deleteTodo}
                    onToggle={toggleComplete}
                    onEdit={openEditModal}
                  />
                ))
              )}
            </AnimatePresence>
          </div>
        </div>
      </motion.div>

      <AnimatePresence>
        {isModalOpen && editingTodo && (
          <EditTodoModal todo={editingTodo} onClose={closeEditModal} onUpdate={updateTodo} />
        )}
      </AnimatePresence>
    </main>
  )
}
