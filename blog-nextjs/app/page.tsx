import Link from "next/link";

export default function Home() {
  return (
    <main>
      <h1> This is a blog App</h1>
      <p>created using Nextjs by ~Kapil Rohilla</p>
      <Link href="/users">
        <button>Goto Users</button>
      </Link>
      <Link href="/posts">
        <button>Goto Posts</button>
      </Link>
    </main>
  );
}
