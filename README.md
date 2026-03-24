# Stack

**A self-funded AI bot that trades and reinvests its profits.**

Stack is an autonomous trading system built to sustain itself. It trades, compounds its own capital, and covers its own costs — without outside funding.

---

## What is Stack?

Stack is an AI-powered trading system with one core constraint: it has to pay for itself.

It starts with a fixed amount of seed capital. From there, it trades autonomously, and every dollar it makes gets reinvested. No money comes in from outside. No profits get pulled out. The system either grows on its own, or it doesn't.

**"Self-funded"** means exactly that — Stack's compute costs, infrastructure, and operation all have to be covered by what it earns from trading. If it can't sustain itself, it fails. That constraint is the whole point.

The goal isn't to build a fund or a product. It's to answer a simple question: can a small, well-designed system compound its way to financial self-sufficiency, built entirely in the open?

---

## The Core Idea

The concept is straightforward:

1. Start with a fixed seed amount
2. Use signals to identify trade opportunities
3. Execute trades with strict risk controls
4. Reinvest all profits back into the portfolio
5. Let compounding do the work over time

The system doesn't need to win every trade. It needs to win consistently enough, over a long enough period, that the curve trends upward. Small edges, repeated often, add up.

The end goal is a system that is fully self-sustaining — paying for its own existence and continuing to grow without any external input.

---

## Why This Exists

Most algorithmic trading projects have the same problems:

- They rely on external capital to keep running
- They only share results when things go well
- They're either fully closed or abandoned open-source repos

Stack is built differently. Everything is documented publicly — the wins, the losses, the strategy changes, the failures. There are no cherry-picked screenshots or curated performance reports.

This is also an experiment in a broader idea: can a software system be designed to sustain itself financially? Not as a product, not as a fund — just as a self-contained system that earns what it needs to keep running.

That question is interesting enough to build toward.

---

## How It Works

```
Market Data → Signal Generation → Risk Management → Execution → Portfolio Update
                                                                        ↓
                                                              Compounding Loop ←
```

**Market Data**
Price feeds and relevant market data are pulled continuously. This is the raw input the system uses to make decisions.

**Signal Generation**
Models analyze the data and surface potential trade opportunities. These signals are based on statistical patterns and learned behavior, not hard-coded rules.

**Risk Management**
Before any order is placed, every trade passes through a risk layer. Position sizing, drawdown limits, and correlation checks determine whether a trade is worth taking.

**Execution**
Qualifying trades are placed automatically through a broker API. Slippage and timing are factored into order logic.

**Portfolio Update**
After each trade, the portfolio is updated and performance is logged. All results feed back into the system.

**Compounding Loop**
Profits are not extracted. They stay in the portfolio and become the base for the next cycle. This loop is the core mechanism Stack is built around.

---

## Core Principles

- **Self-funding** — No external capital. The system earns what it needs to keep running.
- **Compounding over extraction** — All profits are reinvested. Growth comes from repetition, not windfalls.
- **Full transparency** — Everything is built in public. Results are shared as-is.
- **Adaptive learning** — The system is designed to improve over time, not stay static.
- **Risk control** — Protecting capital is treated as more important than chasing returns.

---

## Current Status

| Component | Status |
|---|---|
| Architecture | ✅ Complete |
| Signal framework | ✅ Complete |
| Risk management layer | 🟡 In progress |
| Paper trading | 🟡 In progress |
| Live execution | 🔲 Upcoming |
| Self-sustaining operation | 🔲 Upcoming |

Stack is currently in early development. The core systems are being built and validated in a simulated environment before any live capital is deployed.

No live performance data yet — that comes when it's ready.

---

## Tech Stack

- **Language** — Python 3.11
- **Market data** — Alpaca Markets API, yfinance
- **Signals** — Statistical models, lightweight ML (scikit-learn)
- **Execution** — Alpaca brokerage API
- **Infrastructure** — Single VPS, cron-scheduled jobs
- **Logging** — Local + structured logs, tracked in this repo

Nothing exotic. The goal is a system that's simple enough to understand and maintain, not one that's impressive on paper.

---

## Roadmap

- [x] Initial architecture design
- [x] Data pipeline
- [x] Signal generation framework
- [ ] Risk management layer
- [ ] Paper trading (simulated)
- [ ] Performance validation
- [ ] Live trading deployment
- [ ] Cost coverage milestone (self-funded)
- [ ] Full compounding loop
- [ ] Self-sustaining operation

---

## Built in Public

Everything about Stack is documented openly:

- Strategy decisions and the reasoning behind them
- Code changes and architecture updates
- Real performance numbers — positive and negative
- Post-mortems when things go wrong

The `/log` folder in this repo contains running notes on progress. Issues and Discussions are open. If you see something wrong, or have a question about how something works, open an issue.

No curated results. No highlights reel.

---

## Philosophy

Most trading systems fall into one of two categories.

The first kind is closed. Strategies are proprietary, results are private, and the whole operation runs on external capital. You can read about it, but you can't see it.

The second kind is open source but theoretical. The code is public, but it's never run with real money. There's no skin in the game.

Stack is trying to be a third thing: a real system, running real capital, built entirely in the open. The feedback loop is honest because it has to be. If the system loses money, that's documented. If a strategy doesn't work, it gets replaced. If Stack can't sustain itself, that's a result too.

The long-term goal is survival. Not performance metrics, not an exit — just a system that keeps running and growing on its own, indefinitely.

---

## Disclaimer

Stack is an experimental research project.

Nothing in this repository is financial advice. Trading involves significant risk, including the loss of all capital. Past performance does not predict future results. Do not replicate this system with money you are not prepared to lose entirely.

This project is built for research and learning purposes. It may fail.

---

## License

MIT — use the code however you want, but you're on your own.

---

*This project lives or dies in the open.*
