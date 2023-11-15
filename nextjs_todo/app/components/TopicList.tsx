import React from "react";
import RemoveBtn from "./RemoveBtn";
import Link from "next/link";
import { FiEdit } from "react-icons/fi";
const TodoList = () => {
  return (
    <>
      <div className="p-4 border border-slate-200 my-3 flex justify-between gap-4">
        <div>
          <h2 className="text-2xl">Todo title</h2>
          <div>Todo description</div>
        </div>

        <div className="flex items-center gap-2 ">
          <RemoveBtn />
          <Link href={"/edit/123"}>
            <FiEdit size={24} />
          </Link>
        </div>
      </div>
    </>
  );
};

export default TodoList;
