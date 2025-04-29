'use client';

export default function GuidePage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-sky-500 to-indigo-600 text-white flex flex-col">
      {/* Header */}
      <header className="p-6 flex justify-between items-center">
        <h1 className="text-2xl font-bold tracking-tight">📋 Maless Guide</h1>
        <a
          href="/"
          className="bg-white text-indigo-600 font-semibold px-4 py-2 rounded-lg hover:bg-gray-100 transition"
        >
          Back to Home
        </a>
      </header>

      {/* Main Content */}
      <main className="flex-grow px-6 py-10 flex flex-col items-center">
        <div className="bg-white text-gray-800 rounded-2xl shadow-xl p-8 w-full max-w-3xl">
          <h2 className="text-3xl font-bold mb-6 text-center text-indigo-700">
            📘 Getting Started with Maless App
          </h2>

          <section className="mb-6">
            <h3 className="text-xl font-semibold mb-2">1. What is Maless?</h3>
            <p>
              Maless is a sleek todo-list app that helps you manage your daily tasks efficiently. Stay focused and productive with a simple, intuitive interface.
            </p>
          </section>

          <section className="mb-6">
            <h3 className="text-xl font-semibold mb-2">2. How to Sign In?</h3>
            <ul className="list-disc pl-5 space-y-1">
              <li>Click the <strong>Sign In with Google</strong> button on the homepage.</li>
              <li>Authorize access via your Google account.</li>
              <li>Redirected to your personal task dashboard after Sign In.</li>
            </ul>
          </section>

          <section className="mb-6">
            <h3 className="text-xl font-semibold mb-2">3. Using the Dashboard</h3>
            <ul className="list-disc pl-5 space-y-1">
              <li>View all your tasks organized by tags or status.</li>
              <li>Create new tasks quickly and easily.</li>
              <li>Edit or delete tasks anytime you need.</li>
            </ul>
          </section>

          <section className="mb-6">
            <h3 className="text-xl font-semibold mb-2">4. Tags & Filtering</h3>
            <p>
              Add tags to categorize your tasks. This makes it easier to filter and stay organized by work, personal, or priority levels.
            </p>
          </section>

          <section className="mb-6">
            <h3 className="text-xl font-semibold mb-2">5. Logging Out</h3>
            <p>
              Currently, you can log out by clearing browser cookies. Logout button will be implemented soon.
            </p>
          </section>

          <div className="mt-10 text-center">
            <a
              href="/"
              className="bg-indigo-600 text-white px-6 py-3 rounded-full hover:bg-indigo-700 transition"
            >
              ← Back to Home
            </a>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="text-center text-white/70 text-sm p-4">
        © {new Date().getFullYear()} Maless App. Built with by Dimas Pamungkas & Bobby Pratama.
      </footer>
    </div>
  );
}
