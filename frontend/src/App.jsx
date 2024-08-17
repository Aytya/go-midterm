import './App.css';
import {TodoWrapper} from "./components/TodoWrapper.jsx";
import {PriorityInfo} from "./components/PriorityInfo";

function App() {

    return (
        <div id="App">
            <TodoWrapper />
            <PriorityInfo />
        </div>
    )
}
export default App