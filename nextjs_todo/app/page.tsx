import Image from "next/image";
import TodoList from "./components/TopicList";

export default function Home() {
  return (
    <main>
      <TodoList />
      <TodoList />
      <TodoList />
      <TodoList />
    </main>
  );
}
