"use client";
import { useEffect } from "react";
import { useState } from "react";

let ansArray = []
let k = 0;

export default function Home() {
  let answerObjects = {};
  
  const [questions, setQuestions] = useState([]);
  const [answers, setAnswers] = useState([]);
  const [inputbyuser, SetInputbyuser] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/questions")
      .then((data) => data.json())
      .then((data) => {
        setQuestions(data);
      });
  }, []);

  useEffect(() => {
    fetch("http://localhost:8080/answeredscore")
      .then((data) => data.json())
      .then((data) => {
        setAnswers(data);
      });
  }, []);

  async function createUserAnswer(id, answerbyuser) {
    answerObjects[id] = answerbyuser;
    //ansArray.push(answerObjects)

    const answer = {
      id,
      answerbyuser,
    };
    const result = await fetch("http://localhost:8080/useranswered", {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify(answer)
    }).then((data) => data.json());

    ansArray.push(result)

    //alert(ansArray[k].answerbyuser)
    k += 1
  
  }
  
    function yourScore() {
      let c = 0;
      answers.map((res) => {
        for(let i = 0; i < ansArray.length; i++){
          if (ansArray[i].answerbyuser == res.correctanswer){
            c +=1
          }
        }
      });
      alert("You Score is " + c + " out of " + answers.length)
    }

  return (
    <div>
      <h1>This is a ten question general knowledge quiz. Good Luck!</h1>
      <br/>
      {questions.map((question) => (
        <div class="questions">
          <p id={question.id}>{question.id}. {question.text}</p>
          <div class="checkbox">
            <label>
              <input name={question.id} type="radio" value="1" onClick={() => createUserAnswer(question.id, question.option[0])}/>
              {question.option[0]}</label>
          </div>
          <div class="checkbox">
            <label>
              <input name={question.id} type="radio" value="1" onClick={() => createUserAnswer(question.id, question.option[1])}/>
              {question.option[1]}</label>
          </div>
          <div class="checkbox" >
            <label>
              <input name={question.id} type="radio" value="1" onClick={() => createUserAnswer(question.id, question.option[2])}/>
              {question.option[2]}</label>
          </div>
          <br/>

      </div>
      ))}
          <label>
            <input color="green" type="button" onClick={() => yourScore()}/>
          Submit</label>
    </div>
  )
}