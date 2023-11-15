import React from "react";

const Add = () => {
  return (
    <form className="flex flex-col gap-3">
      <input
        type="text"
        className="border border-slate-500 px-8 py-2 "
        placeholder="Todo Title"
      />
      <input
        type="text"
        className="border border-slate-500 px-8 py-2 "
        placeholder="Todo Description"
      />
      <button className="bg-green-700 py-2 px-6 w-fit text-white font-bold">
        Add Todo
      </button>
    </form>
  );
};

export default Add;
