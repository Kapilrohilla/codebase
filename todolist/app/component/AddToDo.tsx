"use client";
import React, { useState } from "react";
interface dbObj {
  todo: string;
  id: number;
}
const assignId = (previousData: dbObj[]): number => {
  if (previousData.length < 1) {
    return 1;
  }
  return previousData[previousData.length - 1].id + 1;
};
const AddToDo = (props: {
  text: string;
  setText: any;
  db: dbObj[];
  setDb: any;
}) => {
  const handleSubmit = (e: any) => {
    e.preventDefault();
    const newData: dbObj = {
      todo: props.text,
      id: assignId(props.db),
    };
    props.setDb(props.db.concat(newData));

    props.setText("");
  };
  return (
    <form onSubmit={(e) => handleSubmit(e)}>
      <input
        className="border border-black"
        type="text"
        name="addToto"
        id="addTodo"
        placeholder="todo..."
        value={props.text}
        onChange={(e) => props.setText(e.target.value)}
      />
      <button
        className="bg-blue-600 p-1 text-white font-sans border rounded outline-none "
        type="submit"
      >
        Add
      </button>
    </form>
  );
};

export default AddToDo;
