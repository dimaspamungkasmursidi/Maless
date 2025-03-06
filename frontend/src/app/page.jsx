import Image from "next/image";

export default function Home() {
  return (
    <div className="relative flex justify-center items-center min-h-screen">
      <div className="absolute inset-0 bg-gradient-to-br from-slate-900 via-purple-500/70 to-blue-500/80 z-[-10] opacity-50 blur-xl"></div>
      <main className="flex flex-col items-center justify-center">
        <Image
          className="dark:invert mb-12"
          src="/next.svg"
          alt="Next.js logo"
          width={180}
          height={38}
          priority
        />
        <div className="max-w-5xl w-full flex flex-col items-center justify-center font-mono gap-2">
          <h1 className="text-2xl font-bold dark:text-white">
            Welcome to Maless!
          </h1>
          <p className="text-center">Maless is a to-do list app designed to help you stay productive by reducing laziness. The name 'Maless' comes from 'Malas Less,' meaning 'less laziness.' With an intuitive interface and smart task management features, Maless keeps you organized, motivated, and on track to achieve your goals effortlessly!</p>
        </div>
      </main>
    </div>
  );
}
