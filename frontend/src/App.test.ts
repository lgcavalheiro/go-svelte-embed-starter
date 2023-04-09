import { render, screen } from "@testing-library/svelte";
import { describe, expect, test } from "vitest";
import App from "./App.svelte";

describe("App test suite", () => {
  test("Should render", () => {
    render(App);

    expect(screen.getByText("Go + Svelte + Vite")).toBeInTheDocument();
  });
});
