import React from "react";
import { dbObj } from "../page";

function deleteTodo(data: dbObj[], key: number) {
  return data.filter((dbObj) => {
    return dbObj.id !== key;
  });
}
function updateTodo(data: dbObj[], key: number, todo: string) {
  return data.map((dbObj) => {
    if (dbObj.id === key) {
      let newTodo: dbObj = {
        todo,
        id: key,
      };
      return newTodo;
    } else {
      return dbObj;
    }
  });
}
const ListTodo = (props: { data: dbObj[]; setData: any }) => {
  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>S.NO</th>
            <th>Todo</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {props.data.map((obj: dbObj, index, data) => {
            return (
              <tr key={obj.id}>
                <td className="border border-black text-center">{index + 1}</td>
                <td className="border border-black text-justify">{obj.todo}</td>
                <td className=" flex gap-1 border border-black p-1">
                  <button
                    onClick={() => {
                      const newTodo = prompt("Enter string to update")!;
                      if (newTodo) {
                        props.setData(updateTodo(data, obj.id, newTodo));
                      } else {
                        return;
                      }
                    }}
                    className="bg-slate-600 p-0.5 text-white"
                  >
                    Update
                  </button>
                  <button
                    onClick={() => {
                      props.setData(deleteTodo(data, obj.id));
                    }}
                    className=" bg-red-600 p-0.5 text-white"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default ListTodo;
