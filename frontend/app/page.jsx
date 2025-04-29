'use client';

export default function HomePage() {
  const handleLogin = () => {
    window.location.href = 'http://localhost:8081/oauth/auth/google/login';
  };

  const handleGuide = () => {
    window.location.href = '/guide';
  };

  return (
    <div className="min-h-screen flex flex-col justify-between bg-gradient-to-br from-sky-500 to-indigo-600 text-white">
      <header className="p-6 flex justify-between items-center">
        <h1 className="text-2xl font-bold tracking-tight">📋Maless App</h1>
        <button
          onClick={handleGuide}
          className="bg-white text-indigo-600 font-semibold px-5 py-2 rounded-lg hover:bg-gray-100 transition cursor-pointer"
        >
          Guide
        </button>
      </header>

      <main className="flex-grow flex flex-col items-center justify-center text-center px-4">
        <h2 className="text-4xl md:text-5xl font-extrabold mb-6 leading-tight">
          Organize Your Day with <span className="block md:inline">Ease & Clarity</span> <br />
          using Maless App
        </h2>
        <p className="text-lg md:text-xl text-white/80 mb-8 max-w-xl">
          Stay focused, stay productive. Maless helps you plan, track, and complete your tasks effortlessly - one day at a time.
        </p>
        <button
          onClick={handleLogin}
          className="bg-white text-indigo-600 font-bold py-3 px-6 rounded-full shadow-lg hover:bg-gray-100 transition cursor-pointer"
        >
          Get Started - sign in with Google
        </button>
      </main>

      <footer className="text-center text-white/70 text-sm p-4">
        © {new Date().getFullYear()} Maless App. Built with by Dimas Pamungkas & Bobby Pratama.
      </footer>
    </div>
  );
}
