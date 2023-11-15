import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="flex justify-between font-bold bg-slate-800 px-8 py-3 items-center">
      <Link href="/" className=" text-white font-bold">
        Todo List
      </Link>
      <Link href="/add" className="bg-white p-2 rounded-md">
        Add Todo
      </Link>
    </nav>
  );
};

export default Navbar;
