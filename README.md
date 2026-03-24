SELF-FUNDING AI TRADER

A machine that pays for itself - then keeps going.

Self-funded AI trading system that compounds its own capital over time. Built in public.

What This Is
An autonomous trading system that starts with seed capital and is designed to never need outside funding again. Every gain gets rolled back in. Every loss gets studied. The system learns, adapts, and — if everything goes right - grows.
This isn't a hedge fund. It isn't a SaaS product. It's an experiment in whether an AI system can sustain and grow its own financial footprint, documented honestly, in the open.

Core Principles
Self-sufficiency - The system funds its own compute, storage, and operational costs from trading profits. If it can't pay its own bills, it's not working.
Compounding over extraction - Profits stay in. No withdrawals, no performance fees, no exits. The goal is long-term exponential growth, not short-term gains.
Radical transparency - Every strategy decision, loss, win, and code change happens here, in public. No cherry-picked results. No survivorship bias.
Failure is data - Drawdowns get documented. Bad bets get post-mortems. Nothing gets swept under the rug.

How It Works
Market Data → Signal Generation → Risk Filter → Execution → Portfolio Update
                                                                    ↓
                                                          Compounding Loop ←

Data ingestion — Price feeds, order book snapshots, macro indicators
Signal generation — ML models surface trade candidates based on momentum, mean-reversion, and sentiment signals
Risk filter — Position sizing, max drawdown limits, correlation checks before any order is placed
Execution — Automated order placement via broker API with slippage controls
Feedback loop — Live results feed back into model retraining on a rolling basis

Stack

Language — Python 3.11
Data — Alpaca Markets API, yfinance
Models — scikit-learn, lightweight LSTM for sequence signals
Broker API — Alpaca (paper → live)
Infra — Runs on a single VPS, cron-scheduled


Built in Public
Every decision gets documented. Follow along:

Dev log — see /log for running notes on what's working and what isn't
Issues — open an issue if you spot something wrong or have a question
Discussions — strategy ideas, criticism, and questions all welcome

This project lives or dies in the open. If the system blows up its account, that'll be here too.

Philosophy
Most trading systems are either:

Closed-source and secretive
Open-source but abandoned
"Educational" but never run with real money

This is none of those. The code is real. The capital is real. The losses will be real.
The hypothesis: an AI system, given enough time and a sound strategy, can compound a small starting amount into something self-sustaining — covering its own costs and growing beyond them.
We're going to find out.

Disclaimers
This is an experimental research project. Nothing here is financial advice. Past performance doesn't guarantee future results. This project may lose all of its capital. Do not replicate this with money you can't afford to lose.

License
MIT — do whatever you want with the code, but you're on your own.
