import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './App.css'
import {BrowserRouter, Route, Routes} from "react-router";
import Layout from "@/pages/layout.tsx";
import Chat from "@/pages/chat";
import Discover from "@/pages/discover";
import File from "@/pages/file";
import Admin from "@/pages/admin";

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Layout />} >
                    <Route path="/chat" element={<Chat />}/>
                    <Route path="/file" element={<File />}/>
                    <Route path="/discover" element={<Discover />}/>
                    <Route path="/admin" element={<Admin />}/>
                </Route>
            </Routes>
        </BrowserRouter>
    </StrictMode>,
)

