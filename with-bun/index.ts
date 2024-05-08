import { writeFile } from "node:fs";
import { join } from "node:path";

// type Golangify<T> = [T?, Error?]
// function normal(): Golangify<'cheese'> {
//   return ["cheese"];
// }

function normal(callback: (err?: Error, response?: "cheese") => void) {
  if (Math.random() < 0.5) {
    callback(new Error("Oppsioe daisy"), undefined);
  } else {
    callback(undefined, "cheese");
  }
}

normal((err, response) => {
  if (err) {
    throw new Error(JSON.stringify(err));
  }

  console.log(response);
});

// writeFile(
//   join(process.cwd(), "anotherfile.ts"),
//   "console.log(You are poopoo);",
//   { encoding: "utf-8" },
//   (err) => {
//     //
//   }
// );
// import http from "node:http";
// new http.Server((r, res) => {
//   if (r.url === "/") {
//     res.write("Hello world");
//     res.end();
//   }
// }).listen(3001);
// // const obj: Record<string, unknown> = {
// //   name: "alfonso",
// // };

// // function mutatesObj(obj: Record<string, unknown>) {
// //   obj.name = "chicken";
// // }

// // // mutatesObj(obj);
// // mutatesObj(structuredClone(obj));

// // console.log(obj);

// // // Objects and arrays (passed by reference)
// // // primitive values are passed by value
// // // in go: all values are passed by value (copy of the value)
// // // struct -> mutatesObj(struct) -> this would only mutate the copy that is created when the function gets defined
// // // struct -> mutatesObj(&struct) -> this would mutate the actual struct
// // // interfaces -> pointers
