"use client";
import AddToDo from "./component/AddToDo";
import ListTodo from "./component/ListTodo";
import { useState } from "react";
export interface dbObj {
  todo: string;
  id: number;
}
const initialDbData: dbObj[] = [
  {
    id: 1,
    todo: "timepass1",
  },
  { id: 2, todo: "timepass2" },
  {
    id: 3,
    todo: "timepass3",
  },
];
export default function Home() {
  const [db, setDb] = useState(initialDbData);
  const [text, setText] = useState("");
  return (
    <>
      <h1 className="text-center text-3xl">TodoList</h1>
      <main className="flex w-full items-center flex-col mt-5">
        <AddToDo text={text} setText={setText} db={db} setDb={setDb} />
        <ListTodo data={db} setData={setDb} />
      </main>
    </>
  );
}
