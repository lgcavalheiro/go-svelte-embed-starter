import { act, fireEvent, render, screen } from "@testing-library/svelte";
import { describe, expect, test, vi } from "vitest";
import Counter from "./Counter.svelte";

describe("Counter test suite", () => {
  test("Should increase counter on button click", async () => {
    render(Counter);

    expect(screen.getByText("count is 0")).toBeInTheDocument();

    const btn = screen.getByTestId("counter-btn");
    await fireEvent.click(btn);

    expect(screen.getByText("count is 1")).toBeInTheDocument();
  });

  test("Should double counter on button click", async () => {
    render(Counter);

    const counterBtn = screen.getByTestId("counter-btn");
    await fireEvent.click(counterBtn);

    global.fetch = vi.fn().mockImplementationOnce(async (url) => {
      const num = parseInt(url.split("number=")[1], 10);
      return {
        json: () => new Promise((resolve) => resolve({ result: num * 2 })),
      };
    });

    await act(async () => {
      const doubleBtn = screen.getByTestId("double-btn");
      await fireEvent.click(doubleBtn);
    });

    expect(
      screen.getByText("Gopher doubles the counter: 2")
    ).toBeInTheDocument();
  });
});
