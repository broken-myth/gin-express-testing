import express from "express"
const app = express()

app.use(express.json())
app.use(express.urlencoded())


app.get("/get",(req,res)=>{
    res.status(200).send("Hey there we are testing thing out")
})

app.post("/post",(req,res)=>{
    res.status(200).send(`${req.body.username} ${req.body.password}`)
})


app.listen(3000)